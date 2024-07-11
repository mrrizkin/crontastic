import request from "./lib/request";

import Apexchart from "apexcharts";

function buildChart(selector, options = {}) {
  if (typeof selector === "string") {
    selector = document.querySelector(selector);
  }

  if (!selector) {
    throw new Error("Invalid selector");
  }

  const chart = new Apexchart(selector, options);
  chart.render();

  return chart;
}

function spinner(selector, show = true) {
  const element = document.querySelector(selector);
  if (element) {
    element.style = show ? "display: inline-block" : "display: none";
  }
}

function formatBytes(bytes, decimals = 2) {
  if (bytes === 0) return "0 Bytes";

  const k = 1024;
  const dm = decimals < 0 ? 0 : decimals;
  const sizes = ["Bytes", "KB", "MB", "GB", "TB", "PB", "EB", "ZB", "YB"];

  const i = Math.floor(Math.log(bytes) / Math.log(k));
  const value = bytes / Math.pow(k, i);

  const formattedValue = new Intl.NumberFormat("en-US", {
    minimumFractionDigits: dm,
    maximumFractionDigits: dm,
  }).format(value);

  return `${formattedValue} ${sizes[i]}`;
}

document.addEventListener("DOMContentLoaded", function () {
  const halfDonutOptions = {
    chart: {
      type: "donut",
      height: 200,
    },
    plotOptions: {
      pie: {
        startAngle: -90,
        endAngle: 90,
        offsetY: 0,
      },
    },
    grid: {
      padding: {
        bottom: -120,
      },
    },
    series: [0, 100],
    labels: ["Used", "Free"],
    colors: ["#2563eb", "#06b6d4"],
    stroke: {
      colors: ["#2563eb", "#06b6d4"],
    },
    dataLabels: {
      enabled: false,
    },
    legend: {
      show: false,
    },
    tooltip: {
      y: {
        formatter: function (val) {
          return formatBytes(val);
        },
      },
    },
  };

  const diskUsageChart = buildChart("#disk-usage-chart", halfDonutOptions);
  const memoryUsageChart = buildChart("#memory-usage-chart", halfDonutOptions);

  spinner("#disk-usage .loading-spinner", true);
  spinner("#memory-usage .loading-spinner", true);

  request("/api/v1/system_usage")
    .then((response) => {
      const { disk_free, disk_used, mem_free, mem_used } = response.data;
      diskUsageChart.updateSeries([disk_used, disk_free]);
      memoryUsageChart.updateSeries([mem_used, mem_free]);
      spinner("#disk-usage .loading-spinner", false);
      spinner("#memory-usage .loading-spinner", false);
    })
    .catch((error) => {
      console.error(error);
    });
});
