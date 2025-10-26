package main

type CustomWaitGroup struct {
	command chan command
	done    chan struct{}
}

type command struct {
	op  string
	n   int
	rep chan struct{}
}

func NewCustomWaitGroup() *CustomWaitGroup {
	wg := &CustomWaitGroup{
		command: make(chan command),
		done:    make(chan struct{}),
	}

	go func() {
		counter := 0
		waiters := []chan struct{}{}

		for cmd := range wg.command {
			switch cmd.op {
			case "add":
				counter += cmd.n
				if cmd.rep != nil {
					close(cmd.rep)
				}
			case "done":
				counter--
				if counter < 0 {
					// Не паникуем, просто устанавливаем счётчик в 0
					counter = 0
				}
				if cmd.rep != nil {
					close(cmd.rep)
				}
			case "wait":
				if counter == 0 {
					close(cmd.rep)
				} else {
					waiters = append(waiters, cmd.rep)
				}
			}

			if counter == 0 && len(waiters) > 0 {
				for _, ch := range waiters {
					close(ch)
				}
				waiters = nil
			}
		}
		close(wg.done)
	}()

	return wg
}

func (wg *CustomWaitGroup) Add(delta int) {
	rep := make(chan struct{})
	wg.command <- command{op: "add", n: delta, rep: rep}
	<-rep
}

func (wg *CustomWaitGroup) Done() {
	rep := make(chan struct{})
	wg.command <- command{op: "done", n: 1, rep: rep}
	<-rep
}

func (wg *CustomWaitGroup) Wait() {
	rep := make(chan struct{})
	wg.command <- command{op: "wait", rep: rep}
	<-rep
}

func (wg *CustomWaitGroup) Close() {
	close(wg.command)
	<-wg.done
}
