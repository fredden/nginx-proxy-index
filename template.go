<!DOCTYPE html>
<html>
<head>
  <title>nginx-proxy-index</title>
  <meta http-equiv="refresh" content="60">
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
