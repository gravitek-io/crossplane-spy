package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// apiDocs serves a simple HTML page documenting the API endpoints
func apiDocs(c *gin.Context) {
	html := `<!DOCTYPE html>
<html lang="en">
<head>
    <meta charset="UTF-8">
    <meta name="viewport" content="width=device-width, initial-scale=1.0">
    <title>Crossplane Spy API</title>
    <style>
        * { margin: 0; padding: 0; box-sizing: border-box; }
        body {
            font-family: -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            min-height: 100vh;
            padding: 2rem;
        }
        .container {
            max-width: 1200px;
            margin: 0 auto;
            background: white;
            border-radius: 12px;
            box-shadow: 0 20px 60px rgba(0,0,0,0.3);
            overflow: hidden;
        }
        .header {
            background: linear-gradient(135deg, #667eea 0%, #764ba2 100%);
            color: white;
            padding: 3rem 2rem;
            text-align: center;
        }
        .header h1 {
            font-size: 2.5rem;
            margin-bottom: 0.5rem;
        }
        .header p {
            font-size: 1.1rem;
            opacity: 0.9;
        }
        .content {
            padding: 2rem;
        }
        .section {
            margin-bottom: 2rem;
        }
        .section h2 {
            color: #333;
            margin-bottom: 1rem;
            padding-bottom: 0.5rem;
            border-bottom: 2px solid #667eea;
        }
        .endpoint {
            background: #f8f9fa;
            padding: 1rem;
            margin-bottom: 0.5rem;
            border-radius: 6px;
            border-left: 4px solid #667eea;
            transition: all 0.2s;
        }
        .endpoint:hover {
            background: #e9ecef;
            transform: translateX(4px);
        }
        .method {
            display: inline-block;
            padding: 0.25rem 0.75rem;
            background: #28a745;
            color: white;
            border-radius: 4px;
            font-weight: bold;
            font-size: 0.85rem;
            margin-right: 1rem;
        }
        .path {
            font-family: 'Courier New', monospace;
            font-size: 1rem;
            color: #495057;
        }
        .path a {
            color: #667eea;
            text-decoration: none;
        }
        .path a:hover {
            text-decoration: underline;
        }
        .description {
            color: #6c757d;
            margin-top: 0.5rem;
            font-size: 0.9rem;
        }
        .info-box {
            background: #e7f3ff;
            border-left: 4px solid #0066cc;
            padding: 1rem;
            border-radius: 6px;
            margin-bottom: 2rem;
        }
        .info-box strong {
            color: #0066cc;
        }
        .footer {
            text-align: center;
            padding: 2rem;
            color: #6c757d;
            font-size: 0.9rem;
            border-top: 1px solid #dee2e6;
        }
    </style>
</head>
<body>
    <div class="container">
        <div class="header">
            <h1>🔍 Crossplane Spy API</h1>
            <p>REST API for discovering and monitoring Crossplane v2 resources</p>
        </div>

        <div class="content">
            <div class="info-box">
                <strong>ℹ️ API Information</strong><br>
                Base URL: <code>http://localhost:8080</code><br>
                Frontend Dashboard: <a href="http://localhost:3000" target="_blank">http://localhost:3000</a><br>
                All endpoints return JSON data.
            </div>

            <div class="section">
                <h2>System</h2>
                <div class="endpoint">
                    <span class="method">GET</span>
                    <span class="path"><a href="/health" target="_blank">/health</a></span>
                    <div class="description">Health check endpoint</div>
                </div>
            </div>

            <div class="section">
                <h2>Resource Summary</h2>
                <div class="endpoint">
                    <span class="method">GET</span>
                    <span class="path"><a href="/api/v1/resources" target="_blank">/api/v1/resources</a></span>
                    <div class="description">Get a summary count of all Crossplane resources</div>
                </div>
            </div>

            <div class="section">
                <h2>Cluster-Scoped Resources</h2>

                <div class="endpoint">
                    <span class="method">GET</span>
                    <span class="path"><a href="/api/v1/providers" target="_blank">/api/v1/providers</a></span>
                    <div class="description">List all Provider packages</div>
                </div>

                <div class="endpoint">
                    <span class="method">GET</span>
                    <span class="path"><a href="/api/v1/xrds" target="_blank">/api/v1/xrds</a></span>
                    <div class="description">List all Composite Resource Definitions (XRDs)</div>
                </div>

                <div class="endpoint">
                    <span class="method">GET</span>
                    <span class="path"><a href="/api/v1/compositions" target="_blank">/api/v1/compositions</a></span>
                    <div class="description">List all Compositions</div>
                </div>

                <div class="endpoint">
                    <span class="method">GET</span>
                    <span class="path"><a href="/api/v1/functions" target="_blank">/api/v1/functions</a></span>
                    <div class="description">List all Composition Functions</div>
                </div>

                <div class="endpoint">
                    <span class="method">GET</span>
                    <span class="path"><a href="/api/v1/providerconfigs" target="_blank">/api/v1/providerconfigs</a></span>
                    <div class="description">List all ProviderConfigs</div>
                </div>
            </div>

            <div class="section">
                <h2>Composite Resources</h2>

                <div class="endpoint">
                    <span class="method">GET</span>
                    <span class="path"><a href="/api/v1/xrs" target="_blank">/api/v1/xrs</a></span>
                    <div class="description">List all Composite Resource (XR) instances</div>
                </div>
            </div>

            <div class="section">
                <h2>Aggregated Views</h2>

                <div class="endpoint">
                    <span class="method">GET</span>
                    <span class="path"><a href="/api/v1/cluster-resources" target="_blank">/api/v1/cluster-resources</a></span>
                    <div class="description">List all cluster-scoped resources</div>
                </div>

                <div class="endpoint">
                    <span class="method">GET</span>
                    <span class="path"><a href="/api/v1/namespace-resources" target="_blank">/api/v1/namespace-resources</a></span>
                    <div class="description">List all namespace-scoped resources</div>
                </div>
            </div>
        </div>

        <div class="footer">
            <p>Crossplane Spy - Educational tool for Crossplane v2</p>
            <p>Read-only access to Kubernetes resources</p>
        </div>
    </div>
</body>
</html>`

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}
