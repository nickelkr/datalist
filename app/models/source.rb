class Source < ActiveRecord::Base
  validates :url, uniqueness: true

  def self.search(search, page = 1)
    conditions = "%" + search + "%"
    order('name').where('name LIKE ? OR description LIKE ?', conditions, conditions).paginate(page: page, per_page: 10)
  end
end
