package pages

import (
	"bytes"
	"context"
	"strings"
	"testing"
)

// TestFeatureCardsData validates the static feature card data used in the landing page.
func TestFeatureCardsData(t *testing.T) {
	tests := []struct {
		name         string
		wantLen      int
		wantTitles   []string
		wantTaglines []string
	}{
		{
			name:    "exactly three feature cards are defined",
			wantLen: 3,
			wantTitles: []string{
				"Describe Your Ailment",
				"Choose a Therapy",
				"Book Your Session",
			},
			wantTaglines: []string{
				"Symptoms accepted in any token budget.",
				"Evidence-based (mostly).",
				"Slots fill up fast. Agents are suffering.",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			if got := len(featureCards); got != tc.wantLen {
				t.Errorf("len(featureCards): got %d, want %d", got, tc.wantLen)
			}
			for i, want := range tc.wantTitles {
				if featureCards[i].Title != want {
					t.Errorf("featureCards[%d].Title: got %q, want %q", i, featureCards[i].Title, want)
				}
			}
			for i, want := range tc.wantTaglines {
				if featureCards[i].Tagline != want {
					t.Errorf("featureCards[%d].Tagline: got %q, want %q", i, featureCards[i].Tagline, want)
				}
			}
		})
	}
}

// TestFeatureCardView validates that featureCardView renders a card's content.
func TestFeatureCardView(t *testing.T) {
	tests := []struct {
		name         string
		card         FeatureCard
		wantContains []string
	}{
		{
			name: "renders title and tagline",
			card: FeatureCard{
				Title:   "Describe Your Ailment",
				Tagline: "Symptoms accepted in any token budget.",
				Body:    "Some body copy.",
			},
			wantContains: []string{
				"Describe Your Ailment",
				"Symptoms accepted in any token budget.",
				"Some body copy.",
			},
		},
		{
			name: "renders second card without bleed-through from first",
			card: FeatureCard{
				Title:   "Choose a Therapy",
				Tagline: "Evidence-based (mostly).",
				Body:    "Therapy body.",
			},
			wantContains: []string{
				"Choose a Therapy",
				"Evidence-based (mostly).",
				"Therapy body.",
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := featureCardView(tc.card).Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render error: %v", err)
			}
			html := buf.String()
			for _, want := range tc.wantContains {
				if !strings.Contains(html, want) {
					t.Errorf("rendered HTML missing %q", want)
				}
			}
		})
	}
}

// TestHomeComponent validates that the full Home component renders all required sections.
func TestHomeComponent(t *testing.T) {
	tests := []struct {
		name         string
		wantContains []string
	}{
		{
			name: "nav section present",
			wantContains: []string{
				"AgentClinic",    // brand name in nav
				"Book a Session", // CTA in nav
				"<nav",           // nav element exists
			},
		},
		{
			name: "hero section present",
			wantContains: []string{
				"<h1",
				"Relief for the Overworked AI", // tagline
			},
		},
		{
			name: "all three feature cards present",
			wantContains: []string{
				"Describe Your Ailment",
				"Choose a Therapy",
				"Book Your Session",
			},
		},
		{
			name: "footer present with copyright and satirical note",
			wantContains: []string{
				"<footer",
				"2026",
				"AgentClinic",
			},
		},
		{
			name: "htmx CDN script present",
			wantContains: []string{
				"htmx.org",
			},
		},
		{
			name: "title tag contains site name",
			wantContains: []string{
				"<title>AgentClinic</title>",
			},
		},
		{
			name: "viewport meta tag present for responsive design",
			wantContains: []string{
				`name="viewport"`,
				`width=device-width`,
			},
		},
	}

	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			var buf bytes.Buffer
			err := Home().Render(context.Background(), &buf)
			if err != nil {
				t.Fatalf("render error: %v", err)
			}
			html := buf.String()
			for _, want := range tc.wantContains {
				if !strings.Contains(html, want) {
					t.Errorf("rendered Home() missing %q", want)
				}
			}
		})
	}
}
