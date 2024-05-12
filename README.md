# go-dynamic-filters-perf

Since I'm a Go pleb I'm mostly following https://blog.logrocket.com/building-simple-app-go-postgresql/ and adjusted a bit from it. In this repository, I am trying out the difference of performance between a specified query parameters vs. dynamic query parameters in Postgres.

## Test results

### Normal query

#### 10 * 10000

```
Average	
- CPU: 10.21270833
- Memory: 48.63729167

Min
- CPU: 9.54
- Memory: 44.23

Max
- CPU: 11.71
- Memory: 51.53
```