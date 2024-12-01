arkane/
├── cmd/                   # Command-line interface and entry points
│   └── main.go            # Main entry point of the application
├── internal/              # Application-specific internal logic
│   ├── scanner/           # Core scanning functionality
│   │   ├── scanner.go     # Scanning logic for repositories and files
│   │   ├── credentials.go # Credential extraction and handling
│   │   └── orgfinder.go   # GitHub organization discovery
│   ├── github/            # GitHub API wrapper
│   │   ├── client.go      # Client setup and utility methods
│   │   └── repos.go       # Repository-related operations
│   ├── config/            # Configuration management
│   │   └── config.go      # Parsing and loading configuration
│   ├── db/                # Database interactions
│   │   ├── db.go          # General database connection setup
│   │   ├── credentials.go # Methods for storing/retrieving credentials
│   │   └── reports.go     # Methods for storing/retrieving reports
│   └── utils/             # General utilities
│       └── regex.go       # Regex patterns and helpers
├── pkg/                   # Reusable packages
│   ├── logger/            # Logging functionality
│   │   └── logger.go      # Logging utilities
│   └── report/            # Reporting functionality
│       ├── formatter.go   # Formatting scan results
│       └── writer.go      # Writing results to files or other outputs
├── tests/                 # Test files
│   ├── scanner_test.go    # Unit tests for the scanner module
│   ├── db_test.go         # Unit tests for database interactions
│   ├── github_test.go     # Unit tests for GitHub API interactions
│   └── utils_test.go      # Unit tests for utility functions
├── scripts/               # Automation and helper scripts
│   ├── setup.sh           # Script for setting up the environment
│   └── run_tests.sh       # Script for running tests
├── docs/                  # Documentation
│   ├── README.md          # Project description and usage
│   ├── API.md             # API usage and details (if needed)
│   └── CONTRIBUTING.md    # Contribution guidelines
├── .env                   # Environment variables (GitHub tokens, domain, DB credentials, etc.)
├── go.mod                 # Go module file
├── go.sum                 # Go module dependency lock file
└── Makefile               # Build, test, and deploy commands
