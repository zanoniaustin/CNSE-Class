#!/bin/bash
curl -d '{ "voter_id": 1, "name": "Austin", "email": "austin@austin.com" }' -H "Content-Type: application/json" -X POST http://localhost:1080/voters
curl -d '{ "voter_id": 2, "name": "Matt", "email": "matt@matt.com"}' -H "Content-Type: application/json" -X POST http://localhost:1080/voters
curl -d '{ "voter_id": 3, "name": "Jon", "email": "jon@jon.com"}' -H "Content-Type: application/json" -X POST http://localhost:1080/voters
curl -d '{ "voter_id": 4, "name": "Jess", "email": "jess@jess.com"}' -H "Content-Type: application/json" -X POST http://localhost:1080/voters