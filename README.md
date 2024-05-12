# go-dynamic-filters-perf

Since I'm a Go pleb I'm mostly following https://blog.logrocket.com/building-simple-app-go-postgresql/ and adjusted a bit from it. In this repository, I am trying out the difference of performance between a specified query parameters vs. dynamic query parameters in Postgres.

## Test results

### Normal query

#### 10 * 100

```
31.98MiB 
31.98MiB 
31.98MiB 
0.02% 17.66MiB 
0.02% 17.66MiB 
0.01% 17.66MiB 
0.23% 17.69MiB 
0.02% 17.69MiB 
9.73% 44.06MiB 
9.69% 44.13MiB 
9.84% 44.13MiB 
10.32% 44.15MiB 
9.67% 44.7MiB 
10.00% 44.72MiB 
9.72% 44.72MiB 
10.31% 44.71MiB 
10.10% 44.7MiB 
10.18% 44.72MiB 
10.26% 44.72MiB 
9.93% 44.71MiB 
10.01% 43.47MiB 
9.88% 44.11MiB 
9.88% 44.34MiB 
9.85% 44.37MiB 
9.76% 44.94MiB 
9.71% 44.41MiB 
10.39% 44.41MiB 
10.52% 44.42MiB 
10.04% 44.42MiB 
10.00% 44.43MiB 
10.30% 44.33MiB 
9.55% 44.43MiB 
11.04% 44.44MiB 
9.44% 42.86MiB 
9.70% 41.29MiB 
9.49% 38.18MiB 
4.14% 31.88MiB 
0.02% 31.88MiB 

Done! With max response time 2
```