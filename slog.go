// package main

// import (
//     "log/slog"
//     "os"
// )

// func main() {
//  // log to standart output
//  logger:= slog.New(slog.NewTextHandler(os.Stdout,&slog.HandlerOptions{AddSource:true}))

//  // // Log to a specific file
//  // logFile, _ := os.Create("api.log")
//  // logger:= slog.New(slog.NewTextHandler(logFile,&slog.HandlerOptions{AddSource:true}))
//  logger.Info("APP started",slog.String("server","localhost"))
//  // logger.Info("APP started","server","localhost")



//  dbLogger:=logger.With("GROUP","DB") // to GROUP logs
//  dbLogger.Info("DB Cnnnected","Post",270017)
//  dbLogger.Warn("DB Cnnnected2","Post2",270018)

// }