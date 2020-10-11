package utils_test

import (
	"testing"
	"time"

	"metrics-server/pkg/utils"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"k8s.io/component-base/metrics"
)

func TestMetricsUtil(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Prometheus Metrics Utils Test")
}

var _ = Describe("Prometheus Bucket Estimator", func() {
	Context("with a scrape timeout longer than the max default bucket", func() {
		It("should generate buckets in strictly increasing order", func() {
			buckets := utils.BucketsForScrapeDuration(15 * time.Second)
			lastBucket := 0.0
			for _, bucket := range buckets {
				Expect(bucket).To(BeNumerically(">", lastBucket))
				lastBucket = bucket
			}
		})

		It("should include some buckets around the scrape timeout", func() {
			Expect(utils.BucketsForScrapeDuration(15 * time.Second)).To(ContainElement(15.0))
			Expect(utils.BucketsForScrapeDuration(15 * time.Second)).To(ContainElement(30.0))
		})
	})
	Context("with a scrape timeout shorter thant the max default bucket", func() {
		It("should generate buckets in strictly increasing order", func() {
			buckets := utils.BucketsForScrapeDuration(5 * time.Second)
			lastBucket := 0.0
			for _, bucket := range buckets {
				Expect(bucket).To(BeNumerically(">", lastBucket))
				lastBucket = bucket
			}
		})

		It("should include a bucket for the scrape timeout", func() {
			Expect(utils.BucketsForScrapeDuration(5 * time.Second)).To(ContainElement(5.0))
		})
	})
	Context("with a scrape timoeut equalt to the max default bucket", func() {
		maxBucket := metrics.DefBuckets[len(metrics.DefBuckets)-1]
		maxBucketDuration := time.Duration(maxBucket) * time.Second

		It("should generate buckets in strictly inscreasing order", func() {
			buckets := utils.BucketsForScrapeDuration(maxBucketDuration)
			lastBucket := 0.0
			for _, bucket := range buckets {
				Expect(bucket).To(BeNumerically(">", lastBucket))
				lastBucket = bucket
			}
		})

		It("should include a bucket for the scrape timeout", func() {
			Expect(utils.BucketsForScrapeDuration(maxBucketDuration)).To(ContainElement(maxBucket))
		})
	})
})
