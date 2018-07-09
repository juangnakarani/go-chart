package main

import (
	"net/http"
	"time"

	"github.com/juangnakarani/go-chart/drawing"

	"github.com/juangnakarani/go-chart"
)

func drawChart(res http.ResponseWriter, req *http.Request) {
	/*
	   This is an example of using the `TimeSeries` to automatically coerce time.Time values into a continuous xrange.
	   Note: chart.TimeSeries implements `ValueFormatterProvider` and as a result gives the XAxis the appropariate formatter to use for the ticks.
	*/
	ucl := []float64{50.0, 50.0, 50.0, 50.0, 50.0, 50.0, 50.0, 50.0, 50.0, 50.0, 50.0}
	lcl := []float64{-50.0, -50.0, -50.0, -50.0, -50.0, -50.0, -50.0, -50.0, -50.0, -50.0, -50.0}
	wsul := []float64{70.0, 70.0, 70.0, 70.0, 70.0, 70.0, 70.0, 70.0, 70.0, 70.0, 70.0}
	wsll := []float64{-70.0, -70.0, -70.0, -70.0, -70.0, -70.0, -70.0, -70.0, -70.0, -70.0, -70.0}
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true,
			},
		},
		YAxis: chart.YAxis{
			Style: chart.Style{
				Show: true, //enables / displays the y-axis
			},
			Ticks: []chart.Tick{
				{-75.0, "-75"},
				{-50.0, "-50"},
				{-25.0, "-25"},
				{0.0, "0"},
				{25.0, "25"},
				{50.0, "50"},
				{75.0, "75"},
			},
		},

		Background: chart.Style{
			Padding: chart.Box{
				Top:  20,
				Left: 150,
			},
		},
		Series: []chart.Series{
			chart.TimeSeries{
				Name: "Alert Upper Limit",
				XValues: []time.Time{
					time.Now().AddDate(0, 0, -10),
					time.Now().AddDate(0, 0, -9),
					time.Now().AddDate(0, 0, -8),
					time.Now().AddDate(0, 0, -7),
					time.Now().AddDate(0, 0, -6),
					time.Now().AddDate(0, 0, -5),
					time.Now().AddDate(0, 0, -4),
					time.Now().AddDate(0, 0, -3),
					time.Now().AddDate(0, 0, -2),
					time.Now().AddDate(0, 0, -1),
					time.Now(),
				},
				YValues: ucl,
				Style: chart.Style{
					Show:        true,
					StrokeColor: drawing.ColorFromHex("F5B041"),
					StrokeWidth: 2.0,
				},
			},
			chart.TimeSeries{
				Name: "Alert Lower Limit",
				XValues: []time.Time{
					time.Now().AddDate(0, 0, -10),
					time.Now().AddDate(0, 0, -9),
					time.Now().AddDate(0, 0, -8),
					time.Now().AddDate(0, 0, -7),
					time.Now().AddDate(0, 0, -6),
					time.Now().AddDate(0, 0, -5),
					time.Now().AddDate(0, 0, -4),
					time.Now().AddDate(0, 0, -3),
					time.Now().AddDate(0, 0, -2),
					time.Now().AddDate(0, 0, -1),
					time.Now(),
				},
				YValues: lcl,
				Style: chart.Style{
					Show:        true,
					StrokeColor: drawing.ColorFromHex("F5B041"),
					StrokeWidth: 2.0,
				},
			},
			chart.TimeSeries{
				Name: "Work Suspensiun Upper Limit",
				XValues: []time.Time{
					time.Now().AddDate(0, 0, -10),
					time.Now().AddDate(0, 0, -9),
					time.Now().AddDate(0, 0, -8),
					time.Now().AddDate(0, 0, -7),
					time.Now().AddDate(0, 0, -6),
					time.Now().AddDate(0, 0, -5),
					time.Now().AddDate(0, 0, -4),
					time.Now().AddDate(0, 0, -3),
					time.Now().AddDate(0, 0, -2),
					time.Now().AddDate(0, 0, -1),
					time.Now(),
				},
				YValues: wsul,
				Style: chart.Style{
					Show:        true,
					StrokeColor: drawing.ColorFromHex("E74C3C"),
					StrokeWidth: 2.0,
				},
			},
			chart.TimeSeries{
				Name: "Work Suspensiun Lower Limit",
				XValues: []time.Time{
					time.Now().AddDate(0, 0, -10),
					time.Now().AddDate(0, 0, -9),
					time.Now().AddDate(0, 0, -8),
					time.Now().AddDate(0, 0, -7),
					time.Now().AddDate(0, 0, -6),
					time.Now().AddDate(0, 0, -5),
					time.Now().AddDate(0, 0, -4),
					time.Now().AddDate(0, 0, -3),
					time.Now().AddDate(0, 0, -2),
					time.Now().AddDate(0, 0, -1),
					time.Now(),
				},
				YValues: wsll,
				Style: chart.Style{
					Show:        true,
					StrokeColor: drawing.ColorFromHex("E74C3C"),
					StrokeWidth: 2.0,
				},
			},
			chart.TimeSeries{
				Name: "Settlement of Point",
				XValues: []time.Time{
					time.Now().AddDate(0, 0, -10),
					time.Now().AddDate(0, 0, -9),
					time.Now().AddDate(0, 0, -8),
					time.Now().AddDate(0, 0, -7),
					time.Now().AddDate(0, 0, -6),
					time.Now().AddDate(0, 0, -5),
					time.Now().AddDate(0, 0, -4),
					time.Now().AddDate(0, 0, -3),
					time.Now().AddDate(0, 0, -2),
					time.Now().AddDate(0, 0, -1),
					time.Now(),
				},
				YValues: []float64{5.0, 7.0, 6.0, 3.0, 15.0, 9.0, 80.0, 18.0, 19.0, 8.0, 7.0},
				Style: chart.Style{
					Show:        true,
					StrokeColor: drawing.ColorFromHex("3498DB"),
					StrokeWidth: 2.0,
				},
			},
		},
	}

	//note we have to do this as a separate step because we need a reference to graph
	graph.Elements = []chart.Renderable{
		chart.LegendLeft(&graph),
	}

	res.Header().Set("Content-Type", "image/svg+xml")
	graph.Render(chart.SVG, res)
}

func drawCustomChart(res http.ResponseWriter, req *http.Request) {
	/*
	   This is basically the other timeseries example, except we switch to hour intervals and specify a different formatter from default for the xaxis tick labels.
	*/
	graph := chart.Chart{
		XAxis: chart.XAxis{
			Style: chart.Style{
				Show: true,
			},
			ValueFormatter: chart.TimeHourValueFormatter,
		},
		Series: []chart.Series{
			chart.TimeSeries{
				XValues: []time.Time{
					time.Now().Add(-10 * time.Hour),
					time.Now().Add(-9 * time.Hour),
					time.Now().Add(-8 * time.Hour),
					time.Now().Add(-7 * time.Hour),
					time.Now().Add(-6 * time.Hour),
					time.Now().Add(-5 * time.Hour),
					time.Now().Add(-4 * time.Hour),
					time.Now().Add(-3 * time.Hour),
					time.Now().Add(-2 * time.Hour),
					time.Now().Add(-1 * time.Hour),
					time.Now(),
				},
				YValues: []float64{1.0, 2.0, 3.0, 4.0, 5.0, 6.0, 7.0, 8.0, 9.0, 10.0, 11.0},
			},
		},
	}

	res.Header().Set("Content-Type", "image/png")
	graph.Render(chart.PNG, res)
}

func main() {
	http.HandleFunc("/", drawChart)
	http.HandleFunc("/favico.ico", func(res http.ResponseWriter, req *http.Request) {
		res.Write([]byte{})
	})
	http.HandleFunc("/custom", drawCustomChart)
	http.ListenAndServe(":9090", nil)
}
