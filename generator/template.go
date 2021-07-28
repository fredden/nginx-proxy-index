<!DOCTYPE html>
<html>
<head>
  <title>nginx-proxy-index</title>
  <meta http-equiv="refresh" content="60">
  <meta charset="UTF-8" />
  <meta name="viewport" content="width=device-width, initial-scale=1.0" />
  <link rel="preconnect" href="https://fonts.gstatic.com" />
  <link href="https://fonts.googleapis.com/css2?family=Poppins:wght@400;800&display=blocker" rel="stylesheet">
  <style type="text/css">
    body {
      font-family: Poppins, Helvetica, sans-serif;
    }
    h1 {
      font-family: Poppins, Verdana, courier, sans-serif;
    }
    a:hover {
      background-color: #ededed;
    }
    a {
      border-radius: 0.4rem;
      color: inherit;
      display: inline-block;
      height: 1.8rem;
      line-height: 1.8rem;
      padding: 0 0.4rem;
      text-decoration: none;
    }
    @media (prefers-color-scheme: dark) {
      body {
        background-color: #303133;
        color: #ffffff;
      }
      a:hover {
        background-color: #777777;
      }
    }
  </style>
</head>
<body>
  <main>
    <h1>Available Websites</h1>
    <ul>
    {{ range $host, $containers := groupByMulti $ "Env.VIRTUAL_HOST" "," }}
      {{ $host := trim $host }}
      {{ $proto := trim (or (first (groupByKeys $containers "Env.VIRTUAL_PROTO")) "http") }}
      {{ if (ne $host "localhost") }}
      <li><a href="{{ $proto }}://{{ $host }}">{{ $host }}</a></li>
      {{ end }}
    {{ end }}
    </ul>
  </main>
  <!-- https://github.com/fredden/nginx-proxy-index -->
</body>
</html>
