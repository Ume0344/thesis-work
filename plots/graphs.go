package main

import (
	"log"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
	"gonum.org/v1/plot/vg"
)

type data struct {
	Label  string
	Values plotter.Values
}

func main() {
	datumBarChart := []data{
		{
			Label:  "l2fwd",
			Values: plotter.Values{2.7, 2.5, 2.5},
		},
		{
			Label:  "l3-routing",
			Values: plotter.Values{2.6, 3.2, 2.5},
		},
		{
			Label:  "firewall",
			Values: plotter.Values{3.6, 2.6, 1.7},
		},
	}

	datumBarChartTotal := data{
		Label:  "",
		Values: plotter.Values{8, 8.5, 8.3},
	}

	barChart(datumBarChart)
	totalBarChart(datumBarChartTotal)
}

func barChart(datum []data) {
	p := plot.New()

	p.Title.Text = "Provisioning Time For 3 Phases of P4 Deployment"
	p.Y.Label.Text = "Provisioning Time"
	p.Legend.Top = true

	w := vg.Points(20)

	barChart0, err := plotter.NewBarChart(datum[0].Values, w)
	if err != nil {
		log.Fatal(err)
	}
	barChart0.Color = plotutil.Color(0)
	barChart0.LineStyle.Width = vg.Length(0)
	barChart0.Offset = -w

	barChart1, err := plotter.NewBarChart(datum[1].Values, w)
	if err != nil {
		log.Fatal(err)
	}
	barChart1.Color = plotutil.Color(1)
	barChart1.LineStyle.Width = vg.Length(0)

	barChart2, err := plotter.NewBarChart(datum[2].Values, w)
	if err != nil {
		log.Fatal(err)
	}
	barChart2.Color = plotutil.Color(2)
	barChart2.LineStyle.Width = vg.Length(0)
	barChart2.Offset = w

	p.Add(barChart0, barChart1, barChart2)

	p.Legend.Add(datum[0].Label, barChart0)
	p.Legend.Add(datum[1].Label, barChart1)
	p.Legend.Add(datum[2].Label, barChart2)

	p.NominalX("P4-to-C Conversion", "Switch-Compilation", "Switch-Running")
	p.Y.Min = 0
	p.Y.Max = 4

	if err := p.Save(5*vg.Inch, 5*vg.Inch, "barchart.png"); err != nil {
		panic(err)
	}
}

func totalBarChart(datum data) {
	p := plot.New()

	p.Title.Text = "Provisioning Time For End-to-End P4 Deployment"
	p.Y.Label.Text = "Provisioning Time"
	p.Legend.Top = true

	w := vg.Points(20)

	barChart0, err := plotter.NewBarChart(datum.Values, w)
	if err != nil {
		log.Fatal(err)
	}
	barChart0.Color = plotutil.Color(45)
	barChart0.LineStyle.Width = vg.Length(0)
	barChart0.Offset = w / 8

	p.Add(barChart0)

	p.NominalX("l2fwd", "l3-routing", "firewall")
	p.Y.Min = 0
	p.Y.Max = 10

	if err := p.Save(5*vg.Inch, 5*vg.Inch, "totalbarchart.png"); err != nil {
		panic(err)
	}
}
