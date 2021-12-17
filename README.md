# ARGON2 libs

---

# DOCUMENT UNDER DEVELOPMENT

---

While working on [hashkitty](https://github.com/vegasq/hashkitty/) I got curious about alternatives for ARGON2 verification libraries.

## Rules

We run 100000 verifications for:

```$argon2id$v=19$m=16,t=2,p=1$MTExMTExMTE$SzwiO6Uix4CqutzHAncBwQ```

Candidates are `password1`, `password2` .. `password100000`.


| Language    | Lib                      | Time (sec) |
| ----------- | ------------------------ | ---------- |
| Python (C)  | argon2-cffi              | 0.7        |
| Go (C)      | lhecker-argon2           | 0.7        |
| Go (Go)     | matthewhartstonge-argon2 | 1.3        |
| Rust (Rust) | rust-argonautica         | 3          |
| Rust (Rust) | rust-argon2              | 25         |

* All number collected from Macbook Pro 2019, 2.4GHz, 8-Core Intel Core i9

## Conclusion

TBD
