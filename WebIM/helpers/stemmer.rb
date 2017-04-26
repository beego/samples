require 'stemmify'

stemmed_parts = ARGV[0].dup.split(' ').map{ |s| s.stem }.join(' ')
puts stemmed_parts
