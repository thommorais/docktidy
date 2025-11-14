package domain

// DiskUsageRow represents aggregated disk usage metrics for a resource type.
type DiskUsageRow struct {
	Type             string
	Total            int
	Active           int
	SizeBytes        int64
	ReclaimableBytes int64
}

// DiskUsage represents overall Docker disk usage grouped by resource types.
type DiskUsage struct {
	Rows []DiskUsageRow
}
