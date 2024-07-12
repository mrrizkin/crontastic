package handlers

import (
	"syscall"

	"github.com/gofiber/fiber/v2"
	"github.com/pbnjay/memory"

	"github.com/mrrizkin/crontastic/resources/views/pages"
)

type SystemStats struct {
	DiskTotal uint64 `json:"disk_total"`
	DiskFree  uint64 `json:"disk_free"`
	DiskUsed  uint64 `json:"disk_used"`
	MemTotal  uint64 `json:"mem_total"`
	MemFree   uint64 `json:"mem_free"`
	MemUsed   uint64 `json:"mem_used"`
}

func getSystemStats() (*SystemStats, error) {
	var stat syscall.Statfs_t
	err := syscall.Statfs("/", &stat)
	if err != nil {
		return nil, err
	}

	diskTotal := stat.Blocks * uint64(stat.Bsize)
	diskFree := stat.Bfree * uint64(stat.Bsize)
	diskUsed := diskTotal - diskFree

	memoryTotal := memory.TotalMemory()
	memoryFree := memory.FreeMemory()
	memoryUsed := memoryTotal - memoryFree

	return &SystemStats{
		DiskTotal: diskTotal,
		DiskFree:  diskFree,
		DiskUsed:  diskUsed,
		MemTotal:  memoryTotal,
		MemFree:   memoryFree,
		MemUsed:   memoryUsed,
	}, nil
}

func (h *Handlers) SystemUsage(c *fiber.Ctx) error {
	stats, err := getSystemStats()
	if err != nil {
		return h.SendJson(c, fiber.Map{"error": err.Error()}, fiber.StatusInternalServerError)
	}

	return h.SendJson(c, stats, fiber.StatusOK)
}

func (h *Handlers) DashboardPage(c *fiber.Ctx) error {
	return h.Render(c, pages.Dashboard())
}
