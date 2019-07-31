# Documentation

## Index

- [Working](#Working)
  - [Data Storage](#Data-Storage)
  - [Actions](#Actions)
    - [Webhook](#Webhook)
    - [Script](#Script)
- [Configuration](#Configuration)
- [Commands](#Commands)
- [Logging System](#Logging-System)
- [Examples](#Examples)

## Working

_ace_ generates hash of resources using [Page Hash](https://pagehash.muzzammil.xyz) in specified intervals and compares them to previously generated hashes. Incase they don't match, it triggers a specified action and updates the hash stored. This action can be a webhook, a local script, or a program.

To avoid unnecessary triggers, _ace_ tests the URL multiple times to ensure that it is not updated on every single request. By default, _ace_ doesn't add dynamic resources (resources which get updated on every request) but if you wish, you can do so by force with `-f` flag. Let us suppose that you want to add `https://example.com`. _ace_ will generate 5 hashes, consecutively, and compare them. _ace_ then generates a distinct list of those hashes. If there are less than or equal to 2 distinct hashes, _ace_ will add it. If there are more, test will fail and it will not be added. To forcefully add it, you can use `-f` flag like `ace -add -u https://example.com -f`.

### Data Storage

Data files are stored in a directory `.ace` created in user's home directory. Any resource mentioned using `-r` flag will be relative to that `.ace` directory in home. If `-r` points to `ace-test`, path will be generataed as `$HOME/.ace/ace-test`.

### Actions

#### Webhook

When a webhook is set as a trigger (see [configuration](#configuration)), a POST request is made with content type of `application/json`. The JSON sent to webhook is structured as:

```json
{
  "version": "%Version%",
  "data": {
    "url": "%URL%",
    "oldHash": "%Old hash%",
    "newHash": "%New hash%",
    "source": "%Body of updated resource%",
    "lastUpdated": "%Time according to RFC850%",
    "isForced": false
  }
}
```

| Key          | Data Type                               | Description                                                      |
| :----------- | :-------------------------------------- | :--------------------------------------------------------------- |
| Version      | String                                  | Version of ace sending data                                      |
| Data         | Object {String, String, String, String} | Contains information about resource                              |
| URL          | String                                  | URL of changed resource                                          |
| Old Hash     | String                                  | Previously generated hash                                        |
| New Hash     | String                                  | Current hash                                                     |
| Source       | String                                  | Body of changed resource                                         |
| Last Updated | String                                  | When was the resource last changed in format defined by RFC850\* |
| Is forced    | Boolean                                 | Indicates if the resource was forcefully added                   |

#### Script

When a script is set as a trigger (see [configuration](#configuration)), data is passed by command line arguments to the script. You can point it to a binary as well. Arguments are positioned as:

| Argument type | Position | Description                                                      |
| :------------ | :------- | :--------------------------------------------------------------- |
| Source        | First    | Path to temp file where body of changed resource is stored       |
| Old Hash      | Second   | Previously generated hash                                        |
| New Hash      | Third    | Current hash                                                     |
| URL           | Fourth   | URL of changed resource                                          |
| Last Updated  | Fifth    | When was the resource last changed in format defined by RFC850\* |
| Is forced     | Sixth    | Indicates if the resource was forcefully added                   |
| Version       | Seventh  | Version of ace sending data                                      |

## Configuration

_ace_ uses JSON as configuration format which is saved in `.ace-resources` (by default) or a specified data file by using flag `-r` followed by file name.
Structure of JSON is constructed as:

```json
{
  "version": "%Version%",
  "lastRun": {
    "start": "%Time according to RFC850%",
    "end": "%Time according to RFC850%"
  },
  "data": [
    {
      "url": "%URL%",
      "hash": "%HASH%",
      "action": "%{webhook, script}%",
      "location": "%Absolute path%",
      "lastUpdated": "%Time according to RFC850%",
      "isForced": false
    },
    ...
  ]
}
```

| Key          | Data Type               | Description                                                                          |
| :----------- | :---------------------- | :----------------------------------------------------------------------------------- |
| Version      | String                  | Version of ace in which configuration was generated                                  |
| Last Run     | Object {String, String} | Consists of starting and ending time of ace's last run in format defined by RFC850\* |
| Data         | Array                   | Contains an array of information about a resource                                    |
| URL          | String                  | URL of resource including protocol scheme                                            |
| Hash         | String                  | SHA256 generated by Page Hash                                                        |
| Action       | String                  | What to trigger if a resource's hash is changed (can be `webhook` or `script`)       |
| Location     | String                  | Absolute path to webhook or script to trigger                                        |
| Last Updated | String                  | Information about when the resource was last updated in format defined by RFC850\*   |
| Is Forced    | Boolean                 | Indicates if resource was added by force (using `-f`)                                |

## Commands

```text
  -action string
        What to trigger?
        Options: {webhook, script}
        Can be paired with: {-add, -f, -location, -u}
  -add
        Add a resource to crawl
        Can be paired with: {-f, -u, -action, -location}
  -f    Skip test and force add a resource to crawl
        Only to be paired with: {-add}
  -interval float
        Specifies the interval in which ace should crawl in minutes (default 5)
  -location string
        Location of trigger set by action
        Can be paired with: {-action, -add, -f, -u}
        Examples:
         webhook: https://example.com/hook, https://example.com/hook.php
         script: /bin/exec, ~/ace-script.sh
  -r string
        Resource data file
        Can be paired with: {-add, -interval, -remove} (default ".ace-resources")
  -remove
        Remove a resource
        Can be paired with: {-r, -u}
  -u string
        URL
        Can be paired with: {-add, -r, -remove}
  -version
        Print ace version
```

## Logging System

_ace_ write logs to standard output with following format:

```log
Started: %Time according to RFC850%
%Foreach resource%
Checking for %Resource URL% - %{Didn't change/New hash found - %New hash%}%
  Action: %{script/webhook}%
  -- Trigger: %Trigger Location%
  %Details about triggered events%
%End foreach%
Ended: %Time according to RFC850%

--

```

Example:

```log
Started: Monday, 15-Jul-19 00:01:04 IST
Checking for https://example.com - Didn't change
Checking for https://example.com/ToBeUpdatedForScript - New hash found - 3587cb776ce0e4e8237f215800b7dffba0f25865cb84550e87ea8bbac838c423
  Action: script
   -- Trigger: C:\path\to\run.bat
   -- Source saved at: C:\Users\Muzzammil\AppData\Local\Temp\ace-3587cb776ce0e4e8237f215800b7dffba0f25865cb84550e87ea8bbac838c423333130055
   ---Output start---
Data saved in ace-data.txt

   ---Output end---
Checking for https://example.com/ToBeUpdatedForWebhook - New hash found - 3587cb776ce0e4e8237f215800b7dffba0f25865cb84550e87ea8bbac838c423
  Action: webhook
   -- Trigger: https://muzzammil.xyz/ace-hook.php
   -- Response: 200 OK
Ended: Monday, 15-Jul-19 00:01:11 IST

--

...
```

## Examples

### Add a resource

\- using `-add`

```sh
$ ace -add
Enter the URL to add: https://example.com
Testing URL: Test passed
Action - [w]ebhook or [s]cript: webhook
Location: https://muzzammil.xyz/ace-hook.php
https://example.com added to ~/.ace/.ace-resources
```

\- using `-add` `-u` `-action` `-location`

```sh
$ ace -add -u https://example.com -action w -location https://muzzammil.xyz/ace-hook.php
https://example.com added to ~/.ace/.ace-resources
```

### Delete a resource

\- using `-remove`

```sh
$ ace -remove
Enter the URL to remove: https://example.com
https://example.com removed from ~/.ace/.ace-resources
```

\- using `-remove` `-u`

```sh
$ ace -remove -u https://example.com
https://example.com removed from ~/.ace/.ace-resources
```

### Handling webhook (php)

Configuration:

```json
{
  "version": "1.19.7.1",
  "lastRun": {
    "start": "Saturday, 06-Jul-19 23:23:20 IST",
    "end": "Saturday, 06-Jul-19 23:23:22 IST"
  },
  "data": [
    {
      "url": "https://example.com",
      "hash": "1f2409ccb2c81d85a507c2e4e0a4a6c21f8bd282d484476ce9855c88fc9e896b",
      "action": "webhook",
      "location": "https://muzzammil.xyz/ace-hook.php",
      "lastUpdated": "Saturday, 06-Jul-19 23:22:36 IST",
      "isForced": false
    }
  ]
}
```

Code for `ace-hook.php`:

```php
<?php
  $input = file_get_contents("php://input");
  file_put_contents("ace.json", print_r($input, true));
?>
```

Contents of `ace.json` after POST:

```json
{
  "version": "1.19.7.1",
  "data": {
    "url": "https://example.com",
    "oldHash": "1f2409ccb2c81d85a507c2e4e0a4a6c21f8bd282d484476ce9855c88fc9e896b",
    "newHash": "3587cb776ce0e4e8237f215800b7dffba0f25865cb84550e87ea8bbac838c423",
    "source": "<!doctype html>\n<html>\n<head>\n    <title>Example Domain</title>\n\n    <meta charset=\"utf-8\" />\n    <meta http-equiv=\"Content-type\" content=\"text/html; charset=utf-8\" />\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\" />\n    <style type=\"text/css\">\n    body {\n        background-color: #f0f0f2;\n        margin: 0;\n        padding: 0;\n        font-family: \"Open Sans\", \"Helvetica Neue\", Helvetica, Arial, sans-serif;\n        \n    }\n    div {\n        width: 600px;\n        margin: 5em auto;\n        padding: 50px;\n        background-color: #fff;\n        border-radius: 1em;\n    }\n    a:link, a:visited {\n        color: #38488f;\n        text-decoration: none;\n    }\n    @media (max-width: 700px) {\n        body {\n            background-color: #fff;\n        }\n        div {\n            width: auto;\n            margin: 0 auto;\n            border-radius: 0;\n            padding: 1em;\n        }\n    }\n    </style>    \n</head>\n\n<body>\n<div>\n    <h1>Example Domain</h1>\n    <p>This domain is established to be used for illustrative examples in documents. You may use this\n    domain in examples without prior coordination or asking for permission.</p>\n    <p><a href=\"http://www.iana.org/domains/example\">More information...</a></p>\n</div>\n</body>\n</html>\n",
    "lastUpdated": "Saturday, 06-Jul-19 23:22:36 IST",
    "isForced": false
  }
}
```

### Handling script (batch)

Configuration:

```json
{
  "version": "1.19.7.1",
  "lastRun": {
    "start": "Saturday, 13-Jul-19 18:32:29 IST",
    "end": "Saturday, 13-Jul-19 18:32:29 IST"
  },
  "data": [
    {
      "url": "https://example.com",
      "hash": "1f2409ccb2c81d85a507c2e4e0a4a6c21f8bd282d484476ce9855c88fc9e896b",
      "action": "script",
      "location": "C:\\path\\to\\run.bat",
      "lastUpdated": "Saturday, 13-Jul-19 18:32:14 IST",
      "isForced": false
    }
  ]
}
```

Code for `C:\path\to\run.bat`:

```bat
@echo off

echo START > ace-data.txt
echo Source saved at: %1 >> ace-data.txt
echo Old Hash: %2 >> ace-data.txt
echo New Hash: %3 >> ace-data.txt
echo URL: %4 >> ace-data.txt
echo Last updated: %5 >> ace-data.txt
echo Is forced: %6 >> ace-data.txt
echo Version: %7 >> ace-data.txt
echo END >> ace-data.txt

echo Data saved in ace-data.txt
```

Contents of `ace-data.txt`:

```text
START
Source saved at: C:\Users\Muzzammil\AppData\Local\Temp\ace-3587cb776ce0e4e8237f215800b7dffba0f25865cb84550e87ea8bbac838c423084401715
Old Hash: 1f2409ccb2c81d85a507c2e4e0a4a6c21f8bd282d484476ce9855c88fc9e896b
New Hash: 3587cb776ce0e4e8237f215800b7dffba0f25865cb84550e87ea8bbac838c423
URL: https://example.com
Last updated: "Saturday, 13-Jul-19 18:32:14 IST"
Is forced: false
Version: 1.19.7.1
END

```

\* Monday, 02-Jan-06 15:04:05 MST
