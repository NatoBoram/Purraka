# Purraka

[![Build Status](https://travis-ci.org/NatoBoram/Purraka.svg?branch=master)](https://travis-ci.org/NatoBoram/Purraka)
[![Go Report Card](https://goreportcard.com/badge/github.com/NatoBoram/Purraka)](https://goreportcard.com/report/github.com/NatoBoram/Purraka)
[![GoDoc](https://godoc.org/github.com/NatoBoram/Purraka?status.svg)](https://godoc.org/github.com/NatoBoram/Purraka)

Purraka is a market analytics tool. To operate, she needs a MariaDB database. Her goal is to crawl the market and save its content to a database for research purpose.

![Purraka](Images/Full.png)

## What it does

### Crawler

A crawler will periodically access Eldarya's market and copy everything it finds. *Everything.* The database has to be initialized with the bundled `.SQL` file.

### Discord Bot

Eventually, a Discord Bot will come to life and give useful insights about the current state of the market. For now, however, the bot sleeps.

## How to install

1. Run `go get -u -fix github.com/NatoBoram/Purraka`
2. Setup a MariaDB server.
3. Import `database.sql`.
4. Run the bot. During its first run, it will attempt to connect using default credentials and save its configuration in `purreko/database.json`.
5. Edit the file with the appropriate user and password.