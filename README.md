# Shorty
![GitHub repo size](https://img.shields.io/github/repo-size/keivanipchihagh/shorty)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/keivanipchihagh/shorty)
![GitHub License](https://img.shields.io/github/license/keivanipchihagh/shorty)


Make your URLs shorter in Go

## Table of Contents
* [Endpoints](#Endpoints)
* [System Design](#system-design)
* [References](#references)

## Endpoints
| **Path** | **Method** | **Cache** | **Description** 
| --- | --- | -- | -- |
| `/:shortened` | GET | Yes | Redirects to the original URL
| `/urls` | POST | No | Creates a new shortened URL
| `/urls` | GET | No | Returns all URLs available
| `/urls:id` | GET | No | Returns a specific URL by ID
| `/metrics` | GET | No | 	Exposes Prometheus metrics

## System Design

### Hash + Collision Resolution

[Hashing]((https://en.wikipedia.org/wiki/Hash_function)) is a simple and effective method for URL shortening, but it does not guarantee uniqueness. When a collision occurs, we resolve it by adding a [salt](https://en.wikipedia.org/wiki/Salt_(cryptography)) to the hash and rehashing until the collision is resolved.

Pros:
- Fixed URL length
- No need for a unique ID generator

Cons:
- Handling collisions can become resource-intensive
- The process of rehashing can be costly

### Snowflake + Base62

An alternative approach is using Base62 encoding in combination with a unique ID generation system, such as [Twitter's Snowflake](https://github.com/twitter-archive/snowflake) algorithm. This method ensures uniqueness by generating time-based, globally unique IDs.

![Twitter's Snowflake](/assets/snowflake.png)

Pros:
- Easily scalable
- Collision is almost impossible

Cons:
- Requires a unique ID generator
- URL length is not fixed and can increase as IDs grow

## References
- [Hash Function](https://en.wikipedia.org/wiki/Hash_function)
- [Salt (Cryptography)](https://en.wikipedia.org/wiki/Salt_(cryptography))
- [Twitter's Snowflake](https://github.com/twitter-archive/snowflake)
- [System Design Interview An Insider’s Guide](https://www.amazon.com/System-Design-Interview-insiders-Second/dp/B08CMF2CQF) - Chapters 7 & 8