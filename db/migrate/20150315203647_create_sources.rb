class CreateSources < ActiveRecord::Migration
  def change
    create_table :sources do |t|
      t.string :name
      t.string :url
      t.text :description

      t.timestamps null: false
    end
  end
end
