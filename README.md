# Purraka

[![pipeline Status](https://gitlab.com/NatoBoram/Purraka/badges/master/pipeline.svg)](https://gitlab.com/NatoBoram/Purraka/commits/master)
[![Go Report Card](https://goreportcard.com/badge/gitlab.com/NatoBoram/Purraka)](https://goreportcard.com/report/gitlab.com/NatoBoram/Purraka)
[![GoDoc](https://godoc.org/gitlab.com/NatoBoram/Purraka?status.svg)](https://godoc.org/gitlab.com/NatoBoram/Purraka)
[![StackShare](https://img.shields.io/badge/tech-stack-0690fa.svg?style=flat)](https://stackshare.io/NatoBoram/purraka)

Purraka is a market analytics tool. To operate, she needs a MariaDB database. Her goal is to crawl the market and save its content to a database for research purpose.

![Purraka](Images/Full.png)

## What it does

### Crawler

A crawler will periodically access Eldarya's market and copy everything it finds. *Everything.* The database has to be initialized with the bundled `.SQL` file.

### Discord Bot

Purraka will send the cheapest item on the market to a dedicated channel. You can invite her by clicking [here](https://discordapp.com/oauth2/authorize?client_id=426497538263089152&scope=bot&permissions=280576).

#### Commands

Right now, only the basics are here.

##### Setting a callback channel

This is the channel she will be sending items to.

```markdown
@Purraka#4972 set channel callback
```

##### Getting the callback channel

Use this if you're unsure what channel she will send her items to. *If she doesn't respond, make sure she actually has a callback channel.*

```markdown
@Purraka#4972 get channel callback
```

## How to install

1. Run `go get -u -fix gitlab.com/NatoBoram/Purraka`
2. Setup a MariaDB server
3. Import `purraka.sql`
4. Run the bot. During its first run, it will attempt to connect using default credentials and save its configuration in `./purraka/db.json`, `./purraka/header.json` and `./purraka/discord.json`.
5. Edit the file with the appropriate user and password.
