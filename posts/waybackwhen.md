+++
title = "Waybackwhen"
date = "2023-03-03T12:58:32+01:00"
author = ""
authorTwitter = "" #do not include @
cover = ""
tags = ["", ""]
keywords = ["", ""]
description = ""
showFullContent = false
readingTime = false
draft = true
+++

# Slackbot i Rust
Follow these steps https://api.slack.com/authentication/basics
og bruk BlockKit for å lage meldingen du ønsker. Sjekk ut payloaden og bruk den

1. Snakk med Slack med Rust
    - veldig enkel payload, dropper ekstern bibliotek
2. Last ned bilde fra en nettside
    - Headless chrome funket akkurat som forventet
3. Last opp bilde
    - Jeg lurte en periode hvordan jeg skulle gjøre det med bilde fordi Slack
      forventer en lenke. Da må jeg hoste bildet et sted. Fant til slutt
      `uploads`
4. Bruke waybackmachine
https://archive.org/help/wayback_api.php
https://en.wikipedia.org/wiki/Help:Using_the_Wayback_Machine#Specific_archive_copy

Make sure you wait for a specific html tag

5. Lage en liste med nettsider og tidsintervall
Hvordan sjekke om et tidsintervall finnes

Slå alt sammen

Boom

Hvordan skal vi gjøre hosting?
GCP Cloud run , hjelp fra trond


Something when horribly wrong with the environmentvariables. Dont add fnutter.
Work locally, but not in Docker 
This blogpost references this issue. I'm not using docker compose, so it could
be still relevant
https://dev.to/tvanantwerp/don-t-quote-environment-variables-in-docker-268h
https://github.com/docker/compose/issues/8388
Didn't work silently. No idea what the problem was. Most likely I didnt see any
output because the debug_log env variable wasnt set correctly and the slack
token wasnt working. so it did do the screenshot, but failed when uploading the
image. I did see chromium was running but no output. 

I tried strace, but I did see an browser exception in the end. Big timeout
unfortunately. Could be related to me not having logs. Strace didn't help, I
couldnt read the output



# Google cloud jaowsa


Først var planen å ha full CI/CD. Jeg fant en passende action,men den støttet
ikke Google Cloud Run Jobs som jeg ønsket å bruke. Bare service. Derfor bestemte
jeg meg bare for å pushe lokalet til Artifact Registry (ikke container registry
for guds skyld siden det er jo så klart deprecated). Fikk til slutt pushet, og
så satt opp en cloud run jobb. Men når jeg skulle legge inn en secret så
hadde ikke lenger service accounten tilgang til den secreten. Så da ble jeg
usikker, kan jeg ha den i env variabel? Så havnet jeg i et ormehull av IAM
tull. Nå venter jeg på å få riktig tilgang. Se mastodon post:

Pass også på at du har nok RAM. 1 gib gikk, men 512 går ikke. Ellers så virker
Cloud run jobs perfekt til denne oppgaven. 
