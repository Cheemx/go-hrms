package worker

import (
	"context"
	"fmt"
	"log"

	"github.com/Cheemx/go-hrms/internal/config"
)

func SendWeeklyReport(cfg *config.APIConfig) {
	ctx := context.Background()

	summaries, err := cfg.DB.GetWeeklyAttendanceSummary(ctx)
	if err != nil {
		log.Println("failed to fetch weekly attendance summaries:", err)
		return
	}

	fmt.Println(summaries)
}

func SendMonthlyReport(cfg *config.APIConfig) {
	ctx := context.Background()

	summaries, err := cfg.DB.GetMonthlyAttendanceSummary(ctx)
	if err != nil {
		log.Println("failed to fetch weekly attendance summaries:", err)
		return
	}

	fmt.Println(summaries)
}
