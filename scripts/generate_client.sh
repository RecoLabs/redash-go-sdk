#!/bin/bash
cd gen
swagger generate client --name=redashclient --spec=../swagger.yaml --strict-responders
cd ..
