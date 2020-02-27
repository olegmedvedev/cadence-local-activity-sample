# Sample to demonstrate bug in ExecuteLocalActivity

> worker

to start worker

> worker -t

to trigger workflow execution

## there are 2 additional options:

> -l

to start workflow with Local Activity

> -n

activity return nil in result


# Issue is in following case

> worker -t -l -n

When we run workflow with local activity and activity returns nil. In this case workflow wails with result:
```
{
  "reason": "cadenceInternal:Generic",
  "details": "unable to decode argument: 0, **struct {}, with json error: EOF",
  ...
}
```
