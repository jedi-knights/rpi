# Ratings Percentage Index
    
[![GoDoc](https://godoc.org/github.com/levidurfee/rpi?status.svg)](https://godoc.org/github.com/levidurfee/rpi)
[![Go Report Card](https://goreportcard.com/badge/github.com/levidurfee/rpi)](https://goreportcard.com/report/github.com/levidurfee/rpi)
[![Build Status](https://travis-ci.org/levidurfee/rpi.svg?branch=master)](https://travis-ci.org/levidurfee/rpi)
[![Coverage Status](https://coveralls.io/repos/github/levidurfee/rpi/badge.svg?branch=master)](https://coveralls.io/github/levidurfee/rpi?branch=master)
[![License](https://img.shields.io/badge/license-MIT-blue.svg)](

A Go module to provide RPI calculation functionality.

## What Is The RPI

The Rating Percentage Index is a mathematical system for rating sports teams.  The NCAA began developing the RPI in the late 1970s for use in selecting teams to participate in the NCAA Division I Men's Basketball Championship.  The first actual use of the RPI for men's basketball was in 1982.  Over time, the NCAA has expanded use of the RPI to other sports, with the following Division I sports now using it: men's and women's soccer, men's and women's volleyball, women's field hockey, men's and women's ice hockey, men's and women's lacrosse, baseball, softball, and women's water polo.  Interestingly, the NCAA stopped using the RPI for men's basketball beginning with the 2018-19 season, replacing it with the much more complex NET system.  In addition, it now uses that system for women's basketball.  It is not yet known whether the NCAA will make a comparable change for other sports at some point in the future.

The NCAA first used the RPI for Division I women's soccer in 1997.

The way in which the NCAA computes the RPI varies some from sport to sport.  The central structure of the RPI, however, is the same for all sports.

This website deals only with the RPI as used for Division I women's soccer.

## Computing the RPI

The RPI consists of three Elements, plus bonus and penalty adjustments.  It considers only games against Division I opponents.  Below, "Team A" refers to the team whose RPI is being computed.

### Element 1:  Team's Winning Percentage (WP)

WP = (W + 0.5 * T) / (W + L + T)

In this formula, W is Team A's wins; T is Team A's ties; and L is Team A's losses.  Games determined by penalty kicks are considered ties.

So, if Team A has a record of 8 wins, 8 losses, and 4 ties, Element 1 of its RPI is

(8 + (1/2 x 4))/(8 + 8 + 4) = (8 + 2)/20 = 10/20 = .5000

Element 1 tells only Team A's wins and ties compared to its games played.  It tells nothing about the strength of Team A's opponents.  Thus, as an example, Element 1 for a team with an 8-8-4 record against the top 20 Division I teams will be .5000 and Element 1 for a team with an identical record against the bottom 20 Division I teams also will be .5000.

### Element 2: Opponents' Average Winning Percentage Against Other Teams (OWP))

## References

* [RPI: Formula](https://sites.google.com/site/rpifordivisioniwomenssoccer/rpi-formula)