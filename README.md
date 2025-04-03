# Shorty
![GitHub repo size](https://img.shields.io/github/repo-size/keivanipchihagh/shorty)
![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/keivanipchihagh/shorty)
![GitHub License](https://img.shields.io/github/license/keivanipchihagh/shorty)


Make your URLs shorter in Go

* [Endpoints](#Endpoints)
* [System Design](#system-design)

## Endpoints
| **Path** | **Method** | **Cache** | **Description** 
| --- | --- | -- | -- |
| `/:shortened` | GET | Yes | Redirects to the original URL
| `/urls` | POST | No | Creates a new URL
| `/urls` | GET | No | Returns all URLs available
| `/urls:id` | GET | No | Returns a specific URL by ID
| `/metrics` | GET | No | Returns Prometheus metrics

## System Design

### Hash + Collision Resolution

[Hashing]((https://en.wikipedia.org/wiki/Hash_function)) is an easy solution to implement, but it doesn't guarantee unique URLs and the collisions must be solved. In case of a collision, a [salt](https://en.wikipedia.org/wiki/Salt_(cryptography)) is added the the hash and the process is repeated until the collision is resolved.

Pros:
- Has a fixed URL length
- Doesn't need a unique ID generator

Cons:
- Solving collision can become costly

### Snowflake + Base62
Base62 is another approach that requires a unique ID generator to be implemented. To implement the Key Generation Service (KGS), I use [Twitter's Snowflake](https://github.com/twitter-archive/snowflake) algorithm to ensure unique ID generation.

![Twitter's Snowflake](/assets/snowflake.png)

Pros:
- Easily scalable
- Collision is almost impossible

Cons:
- Requires a unique ID generator
- Length is not fixed and grows with IDs

## References
- [Hash Function](https://en.wikipedia.org/wiki/Hash_function)
- [Salt (Cryptography)](https://en.wikipedia.org/wiki/Salt_(cryptography))
- [Twitter's Snowflake](https://github.com/twitter-archive/snowflake)
- [System Design Interview An Insider’s Guide](https://www.amazon.com/System-Design-Interview-insiders-Second/dp/B08CMF2CQF) - Chapters 7 & 8