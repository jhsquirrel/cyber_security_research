#!/bin/sh
export FLASK_APP=xss1.py
export FLASK_ENV=development
python -m flask run --host=0.0.0.0
