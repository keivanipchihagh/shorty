# Shorty
Make your URLs shorter in Go

## System Design
We have two common ways for creating the shortened URLs:

| **Hash + Collision Resolution** | **Base 62 Conversion** |
| --- | --- |
| Has a fixed URL length | Length is not fixed and grows with IDs |
| Hash the URL and add a salt to it | Encode the URL using a base 62 alphabet |
| Doesn't need a unique ID generator | Generated IDs have to be unique |
|Solving collision can become costly | Collision is impossible due to unqiue IDs |

Both methods have their pros and cons. Hashing is an eassy solution to implement, but it doesn't guarantee unique URLs and the collisions must be solved. On the other hand, converting base 62 is a more complex solution that requires a unique ID generator to be implemented.

### Hashing + Collision Resolution

Original URLs are mixed with a [salt](https://en.wikipedia.org/wiki/Salt_(cryptography)) and hashed. The salt further protects against collisions. Then, the hash is computed and the first few characters are used as the shortened URL.

In case of a collision (which is detected when inserting the URL into the database), another salt is added the the hash and the process is repeated until the collision is resolved.

### Key Generation Service (KGS) + Base62 Conversion

The KGS generates a unique ID for each URL. There are multiple ways of doing so:

- **Multi-Master Replication**: Involves having multiple masters at each increament IDs by the number of database servers there are. This approach has some major downsides when scaling the servers.

- **UUID**: Universally unique identifiers are a type of unique identifier that is guaranteed to be unique across all time and space

- **Ticket Server**: A ticket server is responsible for assigning unique and auto-incremented IDs to each URL. It's easy to implement but hard to scale and can be a single point of failure.

- **Twitter's Snowflake**: The idea is to devide the ID into multiple parts instead of directly generating it. It's fairly easy to implement and is highly scalable (for a long time ^_^).

![Twitter's Snowflake](/assets/snowflake.png)


**Final Decision**: I used Twitter's Snowflake (KGS) to generate unique IDs and convert them to base 62 strings.

## References
- [Twitter's Snowflake](https://github.com/twitter-archive/snowflake)
- [System Design Interview An Insider’s Guide](https://www.amazon.com/System-Design-Interview-insiders-Second/dp/B08CMF2CQF) - Chapters 7 & 8
