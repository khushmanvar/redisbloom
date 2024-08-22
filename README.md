# RedisBloom

## Overview

RedisBloom is an advanced implementation of probabilistic data structures in Go, designed to offer high-performance solutions for caching, set membership, and cardinality estimation. This repository features various filters, including a Bloom filter, Counting Bloom filter, Cuckoo filter, and HyperLogLog, all implemented from scratch with a focus on efficiency and scalability.

## Features

- **Bloom Filter**: A space-efficient probabilistic data structure for set membership checks with configurable error rates.
- **Thread-Safe Operations**: Support for concurrent reads and writes, ensuring thread safety in multi-threaded environments.

## Installation

Clone the repository and install dependencies:

```bash
git clone https://github.com/khushmanvar/redisbloom.git
cd redisbloom
go mod tidy