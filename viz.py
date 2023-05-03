#!/usr/bin/env python3

import pandas as pd
import seaborn as sns
import matplotlib.pyplot as plt

from collections import defaultdict



def process(path):
    series = defaultdict(lambda: defaultdict(list))

    with open(path, 'r') as f:
        for line in f:
            parts = line.split()
            mark = parts[0].split("/")
            name = mark[-1].removesuffix("-10")
            size = int(mark[1])
            tput = float(parts[2])

            series[name]["size"].append(size)
            series[name]["tput"].append(tput)

    return series


def reshape(data):
    cols = []
    for name, series in data.items():
        message, serialization = {
            'FlatMarshal': ('flat', 'marshal'),
            'FlatUnmarshal': ('flat', 'unmarshal'),
            'WrappedMarshal': ('wrapped', 'marshal'),
            'WrappedUnmarshal': ('wrapped', 'unmarshal'),
        }[name]

        for x, y in zip(series['size'], series['tput']):
            cols.append([x, y, message, serialization])
    return pd.DataFrame(cols, columns=['event size (bytes)', 'throughput (ns/op)', 'message', 'serialization'])


def viz(data):
    data = reshape(data)
    g = sns.relplot(data=data, x='event size (bytes)', y='throughput (ns/op)', hue='message', style='serialization', palette='muted')
    plt.gca().set_xscale('log', base=2)
    g.set(yscale='log')
    plt.show()


if __name__ == "__main__":
    data = process('wrapper_results.txt')
    viz(data)