# lexir

This repo contains some pieces of codes to play around with words and generate thing such as crosswords, [paraulogic](https://www.vilaweb.cat/paraulogic/) style puzzles, etc ..

## Dictionaries

Download the dictionaries by running

```shell
task download-dictionaries
```

## TODO

1. Download some form a dictionary to local files
2. Decide an interface of common operations to perform in the dictionary
3. Research into multiple implementations of that interface, some possible implementations are:
   1. API-based search
   2. File-based search
   3. Array + binary search
   4. [Trie](https://es.wikipedia.org/wiki/Trie)
   5. [DAFSA](https://en.wikipedia.org/wiki/Deterministic_acyclic_finite_state_automaton)
4. Implement an algorithm to generate the classic crosswords
5. Implement an algorithm to generate [paraulogic](https://www.vilaweb.cat/paraulogic/) style puzzles
6. Create a frontend so people can try to resolve the puzzles.