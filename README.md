# CatBot (Dog Friendly)

CatBot (Dog Friendly) is a Telegram bot with several cat-related (and, by popular demand, dog-related) features:

- /cat: posts a random cat pic
- /catfact: posts a random cat fact
- /dog: posts a random dog pic
- /dogfact: posts a random dog fact

It's publicly available on Telegram as https://t.me/randomcatrobot, but with no guarantees of stability or uptime.


## Configuration

Running the code requires the following environment variables to be set:

- TELEGRAM_TOKEN: Bot API token taken from The Botfather on Telegram
- CAT_FACTS_KEY: API key for the RapidAPI Cat Facts API (see thanks section for URL)


## Docker

You can find Docker containers here: https://hub.docker.com/repository/docker/scottbpc/telegram-catbot

New container images are automatically built from master.

![docker build status](https://img.shields.io/docker/cloud/build/scottbpc/telegram-catbot?style=flat-square)
![docker build automation status](https://img.shields.io/docker/cloud/automated/scottbpc/telegram-catbot?style=flat-square)


## Thanks

Shout out to my wonderful data sources:
- https://thecatapi.com/ -- cat pictures!
- https://rapidapi.com/brianiswu/api/cat-facts -- cat facts!
- https://dog.ceo/ -- dog pictures!
- https://dog-api.kinduff.com/api/facts -- dog facts!
