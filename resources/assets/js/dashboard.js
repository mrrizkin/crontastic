import { buildChart } from "./lib/chart";
import request from "./lib/request";
import { formatBytes, spinner } from "./lib/utils";

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

  request
    .get("/api/v1/system_usage")
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
