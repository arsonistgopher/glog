# glog
Glog: A package to de-couple logging.

This package offers some light de-coupling to logging packages.
Dave Cheney references this issue in his blog post: []()

Whilst this approach couples you to a logging package, it does not do anything fancy.

It offers four methods (Debug/Info/Error/Critical) and does not have any package level vars.

__Simple to Use__
1.  Import a logging package in your main code.
2.  Create a new logging var from a concrete type using your chosen logging package.
3.  Initialise the new logging var with flags and relevant setup.
4.  Instrument your code with glog.Info("message") and other methods exposed.

Done.

See example directory for use examples.

__Driver__

I wrote this because it was a pattern I started to use, so figured instead of replicating, I would build it in to a package.

My logging is now easier.

Fun fact: *Glog* isn't Go log, it's named after it's creator, David Gee (me)
