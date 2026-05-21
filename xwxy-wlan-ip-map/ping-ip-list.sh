#!/bin/sh

COUNT=5

date

sudo ping -c $COUNT  192.168.0.1

sudo ping -c $COUNT  192.168.0.2
sudo ping -c $COUNT  192.168.0.3
sudo ping -c $COUNT  192.168.0.4
sudo ping -c $COUNT  192.168.0.5
sudo ping -c $COUNT  192.168.0.6
sudo ping -c $COUNT  192.168.0.7

sudo ping -c $COUNT  192.168.0.11
sudo ping -c $COUNT  192.168.0.14
sudo ping -c $COUNT  192.168.0.15
sudo ping -c $COUNT  192.168.0.16
sudo ping -c $COUNT  192.168.0.17
sudo ping -c $COUNT  192.168.0.18

date
