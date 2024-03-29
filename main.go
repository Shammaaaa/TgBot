package main

import (
	"log"
	"time"

	tgClient "TgBot/clients/telegram"
	"TgBot/config"
	"TgBot/consumer/event-consumer"
	"TgBot/events/telegram"
	"TgBot/storage/mongo"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "files_storage"
	batchSize   = 100
)

func main() {
	cfg := config.MustLoad()
	//storage := files.New(storagePath)

	storage := mongo.New(cfg.MongoConnectionString, 10*time.Second)

	eventsProcessor := telegram.New(
		tgClient.New(tgBotHost, cfg.TgBotToken),
		storage,
	)

	log.Print("service started")

	consumer := event_consumer.New(eventsProcessor, eventsProcessor, batchSize)

	if err := consumer.Start(); err != nil {
		log.Fatal("service is stopped", err)
	}
}
