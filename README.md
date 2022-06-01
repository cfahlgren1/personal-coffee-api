## Coffee API
---

Trigger the API to turn on the Coffee Maker. The Coffee API server maintains the state of the 508 Coffee Maker. We use Twilio to interact with API via Text Message.

```bash

# Turn on
curl -X POST https://coffee-api.calebfahlgren.com/api/on

# Turn off
curl -X POST https://coffee-api.calebfahlgren.com/api/off
```

Get current state of Coffee Maker

```bash
curl https://coffee-api.calebfahlgren.com/api
```