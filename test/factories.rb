require 'factory_girl'
require 'faker'

FactoryGirl.define do
  sequence :url do |u|
    "http://adatasite#{u}.com"
  end

  factory :source do
    name Faker::Lorem.words()
    url
    description Faker::Lorem.paragraph(10)
  end
end
