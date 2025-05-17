#!/bin/sh

awslocal sns create-topic \
    --name OrderTopic

awslocal sqs create-queue \
    --queue-name ProductionOrderQueue \
