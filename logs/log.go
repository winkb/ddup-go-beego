package logs

func LogOnError(err error, tag string) {
	if err == nil {
		return
	}

	Log().Error(tag+"\n", err.Error())
}
