Q: 我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

A: 我认为应该Wrap这个error，wrap后的error带有堆栈信息可以定位到是哪处的数据库查询代码导致了ErrNoRows, 并且秉承error最多被处理一次和error应该在业务主干逻辑上处理的原则，这个error应该被放到外层去做业务逻辑判断。

