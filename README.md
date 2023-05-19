# Pubsub Hello

https://cloud.google.com/pubsub/docs/emulator

```
gcloud components install pubsub-emulator
gcloud components update
```

## Start

```
gcloud beta emulators pubsub start --project=bocchi-the-rock-0221
```

## Client codes

```env
PUBSUB_EMULATOR_HOST=localhost:8085
PUBSUB_PROJECT_ID=xxx-xxx-111
```
