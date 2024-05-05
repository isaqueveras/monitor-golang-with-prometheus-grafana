package main

import (
	"math/rand"
	"time"

	"github.com/google/uuid"
	"github.com/isaqueveras/outis"
	"github.com/isaqueveras/outis/memory"
	"github.com/prometheus/client_golang/prometheus"
)

var reg = prometheus.NewRegistry()
var metric = NewMetrics(reg)

func Init() {
	watch := outis.Watcher("MyRotine", memory.NewOutis())

	go watch.Go(
		outis.WithID(outis.ID(uuid.NewString())),
		outis.WithName("Routine 01"),
		outis.WithRoutine(func(ctx *outis.Context) {
			now := time.Now()

			time.Sleep(time.Second * time.Duration(rand.Int63n(10)))

			defer func() {
				metric.summary.With(prometheus.Labels{"name": ctx.GetName()}).Observe(time.Since(now).Seconds())
				metric.latency.With(prometheus.Labels{
					"id":   ctx.GetID().ToString(),
					"name": ctx.GetName(),
					"path": ctx.GetPath(),
				}).Observe(time.Since(now).Seconds())
			}()
			metric.info.With(prometheus.Labels{
				"id":   ctx.GetID().ToString(),
				"name": ctx.GetName(),
				"code": uuid.NewString(),
			}).Add(float64(rand.Int()))
		}),
		outis.WithInterval(time.Second*1),
	)

	go watch.Go(
		outis.WithID(outis.ID(uuid.NewString())),
		outis.WithName("Routine 02"),
		outis.WithRoutine(func(ctx *outis.Context) {
			now := time.Now()
			time.Sleep(time.Second * time.Duration(time.Duration(rand.Int63n(10))))
			defer func() {
				metric.summary.With(prometheus.Labels{"name": ctx.GetName()}).Observe(time.Since(now).Seconds())
				metric.latency.With(prometheus.Labels{
					"id":   ctx.GetID().ToString(),
					"name": ctx.GetName(),
					"path": ctx.GetPath(),
				}).Observe(time.Since(now).Seconds())
			}()
			metric.info.With(prometheus.Labels{
				"id":   ctx.GetID().ToString(),
				"name": ctx.GetName(),
				"code": uuid.NewString(),
			}).Add(float64(rand.Int()))
		}),
		outis.WithInterval(time.Second*2),
	)

	go watch.Go(
		outis.WithID(outis.ID(uuid.NewString())),
		outis.WithName("Routine 03"),
		outis.WithRoutine(func(ctx *outis.Context) {
			now := time.Now()
			time.Sleep(time.Second * time.Duration(rand.Int63n(10)))

			defer func() {
				metric.summary.With(prometheus.Labels{"name": ctx.GetName()}).Observe(time.Since(now).Seconds())
				metric.latency.With(prometheus.Labels{
					"id":   ctx.GetID().ToString(),
					"name": ctx.GetName(),
					"path": ctx.GetPath(),
				}).Observe(time.Since(now).Seconds())
			}()
			metric.info.With(prometheus.Labels{
				"id":   ctx.GetID().ToString(),
				"name": ctx.GetName(),
				"code": uuid.NewString(),
			}).Add(float64(rand.Int()))
		}),
		outis.WithInterval(time.Second*3),
	)

	go watch.Go(
		outis.WithID(outis.ID(uuid.NewString())),
		outis.WithName("Routine 04"),
		outis.WithRoutine(func(ctx *outis.Context) {
			now := time.Now()
			time.Sleep(time.Second * time.Duration(rand.Int63n(10)))
			defer func() {
				metric.summary.With(prometheus.Labels{"name": ctx.GetName()}).Observe(time.Since(now).Seconds())
				metric.latency.With(prometheus.Labels{
					"id":   ctx.GetID().ToString(),
					"name": ctx.GetName(),
					"path": ctx.GetPath(),
				}).Observe(time.Since(now).Seconds())
			}()

			metric.info.With(prometheus.Labels{
				"id":   ctx.GetID().ToString(),
				"name": ctx.GetName(),
				"code": uuid.NewString(),
			}).Add(float64(rand.Int()))
		}),
		outis.WithInterval(time.Second*4),
	)

	go watch.Go(
		outis.WithID(outis.ID(uuid.NewString())),
		outis.WithName("Routine 05"),
		outis.WithRoutine(func(ctx *outis.Context) {
			now := time.Now()
			time.Sleep(time.Second * time.Duration(rand.Int63n(10)))
			defer func() {
				metric.summary.With(prometheus.Labels{"name": ctx.GetName()}).Observe(time.Since(now).Seconds())
				metric.latency.With(prometheus.Labels{
					"id":   ctx.GetID().ToString(),
					"name": ctx.GetName(),
					"path": ctx.GetPath(),
				}).Observe(time.Since(now).Seconds())
			}()
			metric.info.With(prometheus.Labels{
				"id":   ctx.GetID().ToString(),
				"name": ctx.GetName(),
				"code": uuid.NewString(),
			}).Add(float64(rand.Int()))
		}),
		outis.WithInterval(time.Second*5),
	)

	go watch.Go(
		outis.WithID(outis.ID(uuid.NewString())),
		outis.WithName("Routine 06"),
		outis.WithRoutine(func(ctx *outis.Context) {
			now := time.Now()

			time.Sleep(time.Second * time.Duration(rand.Int63n(10)))

			defer func() {
				metric.summary.With(prometheus.Labels{"name": ctx.GetName()}).Observe(time.Since(now).Seconds())
				metric.latency.With(prometheus.Labels{
					"id":   ctx.GetID().ToString(),
					"name": ctx.GetName(),
					"path": ctx.GetPath(),
				}).Observe(time.Since(now).Seconds())
			}()
			metric.info.With(prometheus.Labels{
				"id":   ctx.GetID().ToString(),
				"name": ctx.GetName(),
				"code": uuid.NewString(),
			}).Add(float64(rand.Int()))
		}),
		outis.WithInterval(time.Second*6),
	)

	watch.Wait()
}
