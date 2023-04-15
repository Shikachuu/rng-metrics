package cmd

import (
	"math/rand"
	"net/http"
	"os"
	"time"

	"github.com/Shikachuu/podman-example-monitoring/randomMetricsGenerator/pkg"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"golang.org/x/exp/slog"
)

var (
	useJson    bool
	duration   time.Duration
	listenAddr string
)

func RootCommand() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "generate",
		Short: "Start the random metrics generator process",
		RunE: func(cmd *cobra.Command, args []string) error {

			var logHandler slog.Handler

			if useJson {
				logHandler = slog.NewJSONHandler(os.Stdout)
			} else {
				logHandler = slog.NewTextHandler(os.Stdout)
			}

			m := pkg.NewMetrics()
			logger := slog.New(logHandler)

			reg := prometheus.NewRegistry()
			reg.MustRegister(
				collectors.NewGoCollector(),
				collectors.NewProcessCollector(collectors.ProcessCollectorOpts{}),
				collectors.NewBuildInfoCollector(),
				m.DebugCounter,
				m.InfoCounter,
				m.WarnCounter,
				m.ErrorCounter,
			)

			tc := time.Tick(duration)
			go (func() {
				for range tc {
					switch rand.Intn(4) {
					case 0:
						logger.Debug("this is a random log message", "duration", duration, "json", useJson)
						m.DebugCounter.Inc()
					case 1:
						logger.Info("this is a random log message", "duration", duration, "json", useJson)
						m.InfoCounter.Inc()
					case 2:
						logger.Warn("this is a random log message", "duration", duration, "json", useJson)
						m.WarnCounter.Inc()
					case 3:
						logger.Error("this is a random log message", "duration", duration, "json", useJson)
						m.ErrorCounter.Inc()
					}
				}
			})()

			http.Handle("/metrics", promhttp.HandlerFor(reg, promhttp.HandlerOpts{Registry: reg}))

			err := http.ListenAndServe(listenAddr, nil)

			if err != nil {
				return err
			}

			return nil
		},
	}

	cmd.Flags().BoolVar(&useJson, "json", false, "use json as the logging format instead of logfmt")

	cmd.Flags().DurationVarP(&duration, "duration", "d", time.Duration(15000000000), "specify the log and metric generation interval")

	cmd.Flags().StringVarP(&listenAddr, "addr", "a", ":8080", "specify the addres the application should be listening on")

	return cmd
}
