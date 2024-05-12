#!/usr/bin/env bash
while true;
  do docker stats --no-stream --format '{{.CPUPerc}} {{.MemUsage}}' go-dynamic-filters-perf-db-1 | cut -d '/' -f 1 >>docker-stats;
  sleep 1;
done