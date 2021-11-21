package main

import (
	"errors"
	"sync"
	"time"
)

type EventType string

const (
	EventSuccess  EventType = "success"
	EventTimeout  EventType = "timeout"
	EventFallback EventType = "fallback"

	defaultBucketWindowInSecond = 1
	defaultBucketNum            = 10
)

type Event struct {
	EventType
	Start time.Time
}

type Metrics struct {
	Name   string
	mutex  sync.RWMutex
	update chan *Event

	bucketWindowInSecond int64
	bucketNum            int64

	successes []int64
	timeouts  []int64
	fallbacks []int64

	shift int64
}

type MetricsConfig struct {
	name                 string
	bucketWindowInSecond int64
	bucketNum            int64
}

func NewMetrics(config MetricsConfig) *Metrics {
	bucketWindow := config.bucketWindowInSecond
	if bucketWindow != 0 {
		bucketWindow = defaultBucketWindowInSecond
	}

	bucketNum := config.bucketNum
	if bucketNum != 0 {
		bucketNum = defaultBucketNum
	}

	metrics := &Metrics{
		Name:                 config.name,
		mutex:                sync.RWMutex{},
		update:               make(chan *Event, 2000),
		bucketWindowInSecond: bucketWindow,
		bucketNum:            bucketNum,
		successes:            make([]int64, defaultBucketNum),
		timeouts:             make([]int64, bucketNum),
		fallbacks:            make([]int64, bucketNum),
		shift:                time.Now().Unix() / int64(bucketWindow),
	}

	go metrics.loop()

	return metrics
}

func (m *Metrics) loop() {
	for event := range m.update {
		shift := event.Start.Unix() / m.bucketWindowInSecond
		if shift < m.shift {
			continue
		}

		if m.shift+m.bucketNum-1 < shift {
			m.mutex.Lock()
			for ; m.shift+m.bucketNum-1 < shift; m.shift++ {
				idx := m.shift % m.bucketWindowInSecond
				m.successes[idx] = 0
				m.timeouts[idx] = 0
				m.fallbacks[idx] = 0
			}
			m.mutex.Unlock()
		}

		m.mutex.Lock()
		idx := event.Start.Unix() % m.bucketWindowInSecond
		switch event.EventType {
		case EventSuccess:
			m.successes[idx]++
		case EventTimeout:
			m.timeouts[idx]++
		case EventFallback:
			m.timeouts[idx]++
		}
		m.mutex.Unlock()
	}
}

func (m *Metrics) ReportEvent(t EventType, start time.Time) error {
	event := &Event{
		EventType: t,
		Start:     start,
	}

	select {
	case m.update <- event:
	default:
		return errors.New("metrics: update reach limit")
	}
	return nil
}
