package utils

import (
	"fmt"
	"log"
	"time"

	"github.com/getsentry/sentry-go"
)

type Logger struct {
	Debug  bool
	client *sentry.Client
}

func (l *Logger) Init(dsn string, release string) {
	client, err := sentry.NewClient(sentry.ClientOptions{
		Dsn:		dsn,
		Release:	release,
	})
	if err != nil {
		log.Fatalln(err)
	}

	l.client = client
	defer l.client.Flush(5 * time.Second)
}

func (l *Logger) Log(m string) {
	msg := fmt.Sprintf("Log: %v - %v", m, time.Now().Format("2006-01-02T15:04:05-0700"))
	if !l.Debug {
		l.client.CaptureMessage(msg, &sentry.EventHint{}, sentry.NewScope())
	}

	log.Println(msg)
}

func (l *Logger) Error(e error, m string) {
	if !l.Debug {
		l.client.CaptureException(e, &sentry.EventHint{}, sentry.NewScope())
	}

	log.Println(fmt.Sprintf("Fatal: %v - %v", m, time.Now().Format("2006-01-02T15:04:05-0700")), e.Error())
}

func (l *Logger) Fatal(e error, m string) {
	if !l.Debug {
		l.client.CaptureException(e, &sentry.EventHint{}, sentry.NewScope())
		time.Sleep(5 * time.Second)
	}

	log.Fatalln(fmt.Sprintf("Fatal: %v - %v", m, time.Now().Format("2006-01-02T15:04:05-0700")), e.Error())
}
