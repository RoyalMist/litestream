package s3

import "testing"

func TestParseHost(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name               string
		args               args
		wantBucket         string
		wantRegion         string
		wantEndpoint       string
		wantForcePathStyle bool
	}{
		{
			name:               "Hetzner Object Storage",
			args:               args{s: "nbg1.your-objectstorage.com/myBucket"},
			wantBucket:         "myBucket",
			wantRegion:         "nbg1",
			wantEndpoint:       "https://nbg1.your-objectstorage.com",
			wantForcePathStyle: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotBucket, gotRegion, gotEndpoint, gotForcePathStyle := ParseHost(tt.args.s)
			if gotBucket != tt.wantBucket {
				t.Errorf("ParseHost() gotBucket = %v, want %v", gotBucket, tt.wantBucket)
			}
			if gotRegion != tt.wantRegion {
				t.Errorf("ParseHost() gotRegion = %v, want %v", gotRegion, tt.wantRegion)
			}
			if gotEndpoint != tt.wantEndpoint {
				t.Errorf("ParseHost() gotEndpoint = %v, want %v", gotEndpoint, tt.wantEndpoint)
			}
			if gotForcePathStyle != tt.wantForcePathStyle {
				t.Errorf("ParseHost() gotForcePathStyle = %v, want %v", gotForcePathStyle, tt.wantForcePathStyle)
			}
		})
	}
}
