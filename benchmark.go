package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/abadojack/whatlanggo"
	"github.com/pemistahl/lingua-go"
)

type TextSample struct {
	text     string
	language string
}

func main() {
	samples, err := loadTestData("test_samples.txt")
	if err != nil {
		fmt.Printf("Error loading test data: %v\n", err)
		return
	}

	fmt.Printf("Loaded %d test samples\n", len(samples))

	linguaAll := lingua.NewLanguageDetectorBuilder().
		FromAllLanguages().
		Build()

	lingua65 := lingua.NewLanguageDetectorBuilder().
		FromAllSpokenLanguages().
		Build()

	lingua10 := lingua.NewLanguageDetectorBuilder().
		FromLanguages(
			lingua.English,
			lingua.French,
			lingua.German,
			lingua.Spanish,
			lingua.Italian,
			lingua.Portuguese,
			lingua.Dutch,
			lingua.Russian,
			lingua.Chinese,
			lingua.Japanese,
		).
		Build()

	// Benchmark whatlang
	fmt.Println("\n=== whatlang Benchmark ===")
	whatlangResults := benchmarkwhatlang(samples)

	// Benchmark Lingua with all languages
	fmt.Println("\n=== Lingua (All Languages) Benchmark ===")
	linguaAllResults := benchmarkLingua(linguaAll, samples)

	// Benchmark Lingua with 65 spoken languages
	fmt.Println("\n=== Lingua (65 Spoken Languages) Benchmark ===")
	lingua65Results := benchmarkLingua(lingua65, samples)

	// Benchmark Lingua with 10 common languages
	fmt.Println("\n=== Lingua (10 Common Languages) Benchmark ===")
	lingua10Results := benchmarkLingua(lingua10, samples)

	// Print summary
	fmt.Println("\n=== Summary ===")
	fmt.Printf("whatlang: %.2f ms avg, %.2f%% accuracy\n",
		whatlangResults.avgTimeMs,
		whatlangResults.accuracy*100)

	fmt.Printf("Lingua (All): %.2f ms avg, %.2f%% accuracy\n",
		linguaAllResults.avgTimeMs,
		linguaAllResults.accuracy*100)

	fmt.Printf("Lingua (65): %.2f ms avg, %.2f%% accuracy\n",
		lingua65Results.avgTimeMs,
		lingua65Results.accuracy*100)

	fmt.Printf("Lingua (10): %.2f ms avg, %.2f%% accuracy\n",
		lingua10Results.avgTimeMs,
		lingua10Results.accuracy*100)
}

// BenchmarkResult contains the results of a benchmark
type BenchmarkResult struct {
	avgTimeMs float64
	accuracy  float64
}

// benchmarkwhatlang benchmarks the whatlang library
func benchmarkwhatlang(samples []TextSample) BenchmarkResult {
	totalTime := int64(0)
	correct := 0

	for _, sample := range samples {
		start := time.Now()
		detected := whatlanggo.Detect(sample.text)
		elapsed := time.Since(start).Nanoseconds()
		totalTime += elapsed

		// Convert detected language to a string code for comparison
		detectedCode := detected.Lang.Iso6391()

		fmt.Printf("Sample (%s): detected as %s in %d ns\n",
			sample.language,
			detectedCode,
			elapsed)

		if strings.EqualFold(detectedCode, sample.language) {
			correct++
		}
	}

	avgTimeMs := float64(totalTime) / float64(len(samples)) / 1000000.0
	accuracy := float64(correct) / float64(len(samples))

	return BenchmarkResult{
		avgTimeMs: avgTimeMs,
		accuracy:  accuracy,
	}
}

// benchmarkLingua benchmarks the Lingua-Go library
func benchmarkLingua(detector lingua.LanguageDetector, samples []TextSample) BenchmarkResult {
	totalTime := int64(0)
	correct := 0

	for _, sample := range samples {
		start := time.Now()
		detectedLanguage, exists := detector.DetectLanguageOf(sample.text)
		elapsed := time.Since(start).Nanoseconds()
		totalTime += elapsed

		var detectedCode string
		if exists {
			// Convert the Lingua language to ISO 639-1 code
			detectedCode = strings.ToLower(detectedLanguage.IsoCode639_1().String())
		} else {
			detectedCode = "unknown"
		}

		fmt.Printf("Sample (%s): detected as %s in %d ns\n",
			sample.language,
			detectedCode,
			elapsed)

		if strings.EqualFold(detectedCode, sample.language) {
			correct++
		}
	}

	avgTimeMs := float64(totalTime) / float64(len(samples)) / 1000000.0
	accuracy := float64(correct) / float64(len(samples))

	return BenchmarkResult{
		avgTimeMs: avgTimeMs,
		accuracy:  accuracy,
	}
}

// loadTestData loads test samples from a file
func loadTestData(filename string) ([]TextSample, error) {
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		return nil, err
	}

	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var samples []TextSample
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.SplitN(line, "|", 2)
		if len(parts) == 2 {
			samples = append(samples, TextSample{
				language: strings.TrimSpace(parts[0]),
				text:     strings.TrimSpace(parts[1]),
			})
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return samples, nil
}
