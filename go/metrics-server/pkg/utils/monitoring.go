package utils

import (
	"time"

	"k8s.io/component-base/metrics"
)

// BucketsForScrapeDuration calculates a variant of the prometheus default histogram
// buckets that includes relevant buckets around our scrape timeout.
func BucketsForScrapeDuration(scrapeTimeout time.Duration) []float64 {
	// set up some buckets that include our scrape timeout,
	// so that we can easily pinpoint scrape timeout issues.
	// The default buckets provide a sane starting point for
	// the smaller numbers.
	buckets := append([]float64(nil), metrics.DefBuckets...)
	maxBucket := buckets[len(buckets)-1]
	timeoutSeconds := float64(scrapeTimeout) / float64(time.Second)
	if timeoutSeconds > maxBucket {
		// [defaults, (scrapeTimeout + (scrapeTimeout - maxBucket)/ 2), scrapeTimeout, scrapeTimeout*1.5, scrapeTimeout*2]
		halfwayToScrapeTimeout := maxBucket + (timeoutSeconds-maxBucket)/2
		buckets = append(buckets, halfwayToScrapeTimeout, timeoutSeconds, timeoutSeconds*1.5, timeoutSeconds*2)
	} else if timeoutSeconds < maxBucket {
		var i int
		var bucket float64
		for i, bucket = range buckets {
			if bucket > timeoutSeconds {
				break
			}
		}

		if bucket-timeoutSeconds < buckets[0] || (i > 0 && timeoutSeconds-buckets[i-1] < buckets[0]) {
			// if we're sufficiently close to another bucket, just skip this
			return buckets
		}

		// likely that our scrape timeout is close to another bucket, so don't bother injecting more that just our scrape timeout
		oldRest := append([]float64(nil), buckets[i:]...) // make a copy so we don't overwrite it
		buckets = append(buckets[:i], timeoutSeconds)
		buckets = append(buckets, oldRest...)
	}

	return buckets
}
