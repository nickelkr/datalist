class Source < ActiveRecord::Base
  validates :url, uniqueness: true
end
