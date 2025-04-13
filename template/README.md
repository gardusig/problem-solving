## Compile options

### C++

```shell
g++-14 \
    -std=c++23 \
    -O3 \
    -Wall \
    -Wl,-stack_size -Wl,10000000 \
    main.cpp
```
