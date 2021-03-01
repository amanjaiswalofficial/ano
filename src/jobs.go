package src

import "github.com/sirupsen/logrus"

// N is used by the counter server
var N int = 1

// Init logging
func init() {
	logrus.SetLevel(logrus.InfoLevel)
	logrus.SetFormatter(&logrus.TextFormatter{FullTimestamp: true})
}

// Count increments N by 1 at given point in time
func Count() {
	logrus.Infof("Increment N by 1 => %d", N)
	N++
}
