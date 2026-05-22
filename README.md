<h1 align="center"> key-value style in-memory database. </h1>
<br>
<ul>
  <li>persistant input/error handling for edge cases.</li>
  <li>interactive cli mode, which stores variable-like data (for now only supports values of type string and int) to then be retrieved         later with methods like "GET", "SET", "DEL", and "EXIT" to exit the program.</li>
  <li>serializes values to bytes before appending to the database struct for optimised performance.</li>
</ul>
<br>
<h1 align="center">interactive cli tutorial:</h1>
<h2 align="center">run the following to begin:</h2>
<h3 align="center">
  <code> go run . </code>
</h3>
<h2 align="center">and follow up by declaring a variable:</h2>
<h3 align="center">
  <code> SET [KeyName] [Value] </code>
</h3>
<h3 align="center">(key names only support either integers or strings as of now).</h3>

<h2 align="center">after setting a variable, you can then retrieve it like so:</h2>
<h3 align="center">
  <code> GET [KeyName] </code>
</h3>
<h3 align="center">and it returns the value of the key given.</h3>

<h2 align="center">to delete a variable/key:</h2>
<h3 align="center">
  <code> DEL [KeyName] </code>
</h3>

<h2 align="center">finally, to exit the program, run:</h2>
<h3 align="center">
  <code> EXIT </code>
</h3>
