import Apexchart from "apexcharts";

export function buildChart(selector, options = {}) {
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
