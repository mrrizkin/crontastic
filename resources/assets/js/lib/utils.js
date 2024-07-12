export function spinner(selector, show = true) {
  const element = document.querySelector(selector);
  if (element) {
    element.style = show ? "display: inline-block" : "display: none";
  }
}

export function formatBytes(bytes, decimals = 2) {
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
