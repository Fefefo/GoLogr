# GoLogr

## Description
A CLI tool written in Go to analyze log files and extract useful insight such as log level distribution and most frequent errors.

## Demo
![Alt Text](../assets/demo.gif?raw=true)

## Installation
git clone https://github.com/Fefefo/GoLogr.git
cd GoLogr
go build -o GoLogr

## Usage
### Basic usage
./GoLogr \--file test.log
### Filter by log level
./GoLogr \--file test.log \--level ERROR
### Show top N errors
./GoLogr \--file test.log \--top 3
### JSON output
./GoLogr \--file test.log \--json
